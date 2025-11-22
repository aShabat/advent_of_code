(ns aoc
  (:require [clojure.string :refer [upper-case]]
            [clojure.java.shell :refer [sh]]))

(def aoc-session "53616c7465645f5f009a193d45549a164835a9d33341fd84a2923f27294073dfb82735276f3a6e50b46478366b73999f20a83f31c0e5303ddb280b96bdc7e658")
(def http-client (java.net.http.HttpClient/newHttpClient))

(defn http-request
  ([url method body]
   (let [body-object (if (nil? body)
                       (java.net.http.HttpRequest$BodyPublishers/noBody)
                       (java.net.http.HttpRequest$BodyPublishers/ofString body))]
     (prn body-object)
     (-> (java.net.http.HttpRequest/newBuilder)
         (.uri (java.net.URI. url))
         (.method (upper-case method) body-object)
         (.header "Cookie" (str "session=" aoc-session))
         (.build))))
  ([url method body & headers]
   (let [body-object (if (nil? body)
                       (java.net.http.HttpRequest$BodyPublishers/noBody)
                       (java.net.http.HttpRequest$BodyPublishers/ofString body))]
     (prn body-object)
     (-> (java.net.http.HttpRequest/newBuilder)
         (.uri (java.net.URI. url))
         (.method method body-object)
         (.headers (into-array headers))
         (.build)))))

(defn http-send
  [request]
  (let [response (.send http-client request (java.net.http.HttpResponse$BodyHandlers/ofString))]
    {:body (.body response) :status (.statusCode response)}))

(defn aoc-request [method body & args]
  (let [aoc-url (apply str "https://adventofcode.com" (map #(str "/" %) args))]
    (http-request aoc-url method body)))

(defn html-extract-main [html] (get (re-find #"(?s)<main>(.*)</main>" html) 1))

(defn aoc-get-exercise-html [year day]
  (let [request (aoc-request "get" nil year "day" day)
        response (http-send request)]
    (html-extract-main (:body response))))

(defn convert-file [file from to] (sh "pandoc" "-o" (str file to) (str file from)))

(defn aoc-get-exercise [year day]
  (let [html (aoc-get-exercise-html year day)
        file (str "static/" day)]
    (spit  (str file ".html") html)
    (convert-file file ".html" ".md")))

(aoc-get-exercise 2021 1)
