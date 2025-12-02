(ns day2
  (:require [clojure.string :as string]))
(load-file "util.clj")

(def test-input "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")
(def full-input (util/read-input 2))

(defn parse-input [input]
  (->> (string/split input #",")
       (map #(map parse-long (string/split % #"-")))))

(def pow10 (memoize (fn [n]
                      (if (= n 0) 1
                          (if (even? n) (let [root (pow10 (/ n 2))] (* root root))
                              (* 10 (pow10 (dec n))))))))

(apply max (map #(- (second %) (first %)) (parse-input full-input)))

(defn dumb-is-invalid-part-1 [id]
  (->> id
       (str)
       (re-matches #"(\d+)\1")))

(defn dumb-part-1 [input]
  (->> input
       (parse-input)
       (map #(range (first %) (inc (second %))))
       (map #(filter dumb-is-invalid-part-1 %))
       (map #(apply + %))
       (apply +)))

(dumb-part-1 test-input) ;should be 1227775554

(util/aoc-send-answer 2 1 (dumb-part-1 full-input))

(util/aoc-get-exercise 2)

(defn dumb-is-invalid-part-2 [id]
  (->> id
       (str)
       (re-matches #"(\d+)\1+")))

(defn dumb-part-2 [input]
  (->> input
       (parse-input)
       (map #(range (first %) (inc (second %))))
       (map #(filter dumb-is-invalid-part-2 %))
       (map #(apply + %))
       (apply +)))

(dumb-part-2 test-input)

(util/aoc-send-answer 2 2 (dumb-part-2 full-input))
