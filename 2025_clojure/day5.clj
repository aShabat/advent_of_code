(ns day5
  (:require [clojure.string :refer [split]]))
(load-file "util.clj")

(def test-input "3-5
10-14
16-20
12-18

1
5
8
11
17
32")

(def real-input (util/read-input 5))

(defn parse-input
  [input]
  (let [input (util/lines input)
        empty-line (first (filter #(= "" (input %)) (range (count input))))]
    [(->> (take empty-line input)
          (mapv #(split % #"-"))
          (mapv (partial mapv parse-long)))
     (->> (drop (inc empty-line) input)
          (map parse-long))]))

(defn collapse-reducer
  [[acc last-interval] new-interval]
  (if (>= (last-interval 1) (new-interval 0))
    [acc [(last-interval 0) (max (last-interval 1) (new-interval 1))]]
    [(conj acc last-interval) new-interval]))

(defn collapse-intervals
  [intervals]
  (subvec (->> intervals
               (sort #(apply compare (map first %&)))
               (reduce collapse-reducer [[] [0 0]])
               (#(conj (% 0) (% 1))))
          1))

(defn in-interval?
  [number intervals]
  (if (= (count intervals) 1)
    (and (>= number (get-in intervals [0 0]))
         (<= number (get-in intervals [0 1])))
    (let [middle (quot (count intervals) 2)]
      (if (>= number (get-in intervals [middle 0]))
        (recur number (subvec intervals middle))
        (recur number (subvec intervals 0 middle))))))

(defn part-1
  [input]
  (let [[intervals numbers] (parse-input input)
        intervals (collapse-intervals intervals)]
    (count (filter #(in-interval? % intervals) numbers))))

(util/aoc-send-answer 5 1 (part-1 real-input))

(defn part-2
  [input]
  (->> input
       parse-input
       first
       collapse-intervals
       (map #(inc (- (% 1) (% 0))))
       (reduce +)))

(util/aoc-send-answer 5 2 (part-2 real-input))
