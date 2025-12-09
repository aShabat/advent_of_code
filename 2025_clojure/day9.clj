(ns day9
  (:require [clojure.string :refer [split]]))
(load-file "util.clj")

(def test-input "7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3")

(def real-input (util/read-input 9))

(defn parse-input
  [input]
  (mapv #(mapv parse-long (split % #",")) (util/lines input)))

(defn area [[x1 y1] [x2 y2]] (* (inc (abs (- x1 x2))) (inc (abs (- y1 y2)))))

(defn part-1
  [input]
  (let [points (parse-input input)]
    (apply max (for [p1 points p2 points] (area p1 p2)))))

; (util/aoc-send-answer 9 1 (part-1 real-input))

(defn rectangle-intercects-line
  [[[x1 y1] [x2 y2]] [[x3 y3] [x4 y4]]]
  (if (= x3 x4)
    (recur [[y1 x1] [y2 x2]] [[y3 x3] [y4 x4]])
    (and (> y3 (min y1 y2))
         (< y3 (max y1 y2))
         (<= (min x3 x4) (max x1 x2))
         (>= (max x3 x4) (min x1 x2)))))

(defn rectangle-in-bound
  [rectangle points]
  (let [lines
        (map (comp vec list) points (conj (subvec points 1) (first points)))]
    (not (some (partial rectangle-intercects-line rectangle) lines))))

(defn part-2
  [input]
  (let [points (parse-input input)
        rectangles (for [p1 points p2 points] [p1 p2])]
    (->> rectangles
         (filter #(rectangle-in-bound % points))
         (map #(apply area %))
         (apply max))))

; (util/aoc-send-answer 9 2 (part-2 real-input))
