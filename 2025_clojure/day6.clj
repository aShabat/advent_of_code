(ns day6
  (:require [clojure.string :refer [split]]))
(load-file "util.clj")

(def test-input
  "123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  ")

(def real-input (slurp "static/6/input.txt"))

(defn parse-word
  [word]
  (case word
    "+" +
    "*" *
    (parse-long word)))

(defn parse-input-1
  [input]
  (->> input
       util/lines
       (map #(split % #"\s+"))
       (map (partial filter #(not= "" %)))
       (mapv (partial mapv parse-word))))


(defn part-1
  [input]
  (->> (parse-input-1 input)
       reverse
       (apply mapv list)
       (map eval)
       (reduce +)))

(part-1 real-input)

(defn parse-column
  [[number-chars [function?]]]
  [(->> number-chars
        (filter #(not= \space %))
        (apply str)
        parse-long)
   (case function?
     \+ +
     \* *
     nil)])

(defn parse-input-2
  [input]
  (->> input
       (#(split % #"\n"))
       (apply mapv (comp vec list))
       (map #(split-at (dec (count %)) %))
       (map parse-column)))

(defn part-2
  [input]
  (->> input
       parse-input-2
       (reduce (fn [[sum fun acc :as r] [number function? :as n]]
                 (prn r n)
                 (cond (nil? fun) [sum function? (bigint number)]
                       (nil? number) [(+ sum acc) nil nil]
                       :default [sum fun (fun acc number)]))
               [0N nil nil])
       (#(+ (% 0) (% 2)))))

(parse-input-2 real-input)
(part-2 real-input)
