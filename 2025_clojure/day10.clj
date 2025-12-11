(ns day10
  (:require [clojure.string :refer [trim split]]))
(load-file "util.clj")

(def test-input
  "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
    [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
    [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}")

(def real-input (util/read-input 10))

(defn parse-lights
  [lights]
  (reduce #(+ (* 2 %1) (if (= %2 \#) 1 0)) 0 (reverse lights)))

(defn parse-switch [bits] (reduce #(bit-set %1 %2) 0 bits))

(defn parse-input
  [input]
  (map (fn [line]
         (let [matches (re-seq #"[\[\(\{](\S*)[\]\)\}]" line)
               matches (mapv #(% 1) matches)]
           {:lights (parse-lights (matches 0)),
            :bin-switches (set (map #(parse-switch (mapv parse-long
                                                         (split % #",")))
                                    (subvec matches 1 (dec (count matches))))),
            :switches (set (map #(mapv parse-long (split % #","))
                                (subvec matches 1 (dec (count matches))))),
            :joltage (mapv parse-long
                           (split (matches (dec (count matches))) #","))}))
       (map trim (util/lines input))))

(defn power
  [[f & r]]
  (if f
    (let [power-r (power r)] (into power-r (map #(conj % f) power-r)))
    [[]]))

(defn part-1
  [input]
  (->> input
       parse-input
       (map (fn [{:keys [lights bin-switches]}]
              (->> bin-switches
                   power
                   (filter #(= lights (apply bit-xor 0 0 %)))
                   (sort-by count)
                   first
                   count)))
       (reduce +)))

; (util/aoc-send-answer 10 1 (part-1 real-input))
