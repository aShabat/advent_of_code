(ns util
  (:require [clojure.string :refer [upper-case split trim]]
            [clojure.java.shell :refer [sh]]
            [clojure.java.io :refer [file]]))

(defn echo [x] (prn (str "ECHO: " x)) x)

(def aoc-session
  "53616c7465645f5f03a5fb622566637eec99981758671569dbd68fac67fe178c1a2f68f93f63176d89cfc6770c6acce2260addfb6f5143d1a2c33d24adb9e5c8")
(def year 2025)
(def http-client (java.net.http.HttpClient/newHttpClient))

(defn http-request
  ([url method body]
   (let [body-object (if (nil? body)
                       (java.net.http.HttpRequest$BodyPublishers/noBody)
                       (java.net.http.HttpRequest$BodyPublishers/ofString
                        body))]
     (->
       (java.net.http.HttpRequest/newBuilder)
       (.uri (java.net.URI. url))
       (.method (upper-case method) body-object)
       (.header "Cookie" (str "session=" aoc-session))
       (.header
        "User-Agent"
        "https://github.com/aShabat/advent_of_code/blob/main/2025_clojure/aoc.clj")
       (.header "Content-Type" "application/x-www-form-urlencoded")
       (.build)))))

(defn- http-send
  [request]
  (let [response (.send http-client
                        request
                        (java.net.http.HttpResponse$BodyHandlers/ofString))]
    {:body (.body response), :status (.statusCode response)}))

(defn- aoc-request
  [method body & args]
  (let [aoc-url (apply str "https://adventofcode.com" (map #(str "/" %) args))]
    (http-request aoc-url method body)))

(defn- html-extract-main [html] (get (re-find #"(?s)<main>(.*)</main>" html) 1))

(defn- dir-exists?
  [path]
  (let [f (file path)] (when-not (.isDirectory f) (.mkdirs f))))

(defn- aoc-get-exercise-html
  [day]
  (let [request (aoc-request "get" nil year "day" day)
        response (http-send request)]
    (html-extract-main (:body response))))

(defn- convert-file
  [file from to]
  (sh "pandoc" "-o" (str file to) (str file from)))

(defn aoc-get-exercise
  [day]
  (let [html (aoc-get-exercise-html day)
        dir (str "static/" day)
        file (str dir "/exercise")]
    (dir-exists? dir)
    (spit (str file ".html") html)
    (convert-file file ".html" ".md")))

(defn- aoc-send-answer-html
  [day part answer]
  (let [body (str "level=" part "&answer=" answer)
        request (aoc-request "post" body year "day" day "answer")
        response (http-send request)]
    (html-extract-main (:body response))))

(defn aoc-send-answer
  [day part answer]
  (let [html (aoc-send-answer-html day part answer)]
    (condp re-find html
      #"(?s)That's the right answer" (do (prn "Right answer")
                                         (aoc-get-exercise day))
      #"(?s)That's not the right answer" :wrong
      #"(?s)You gave an answer too recently" :too-soon)))

(defn aoc-get-input
  [day]
  (let [request (aoc-request "get" nil year "day" day "input")
        response (http-send request)
        dir (str "static/" day)
        file (str dir "/input.txt")]
    (dir-exists? dir)
    (spit file (:body response))))

(defn read-input
  [day]
  (let [file (str "static/" day "/input.txt")] (trim (slurp file))))

(defn lines [input] (split input #"\n"))

(defmacro defnm [name params & body] `(def ~name (memoize (fn ~params ~@body))))

(defn cartesian
  ([] '())
  ([x] (map list x))
  ([x & ys]
   (for [element-x x
         element-y (apply cartesian ys)]
     (cons element-x element-y))))
