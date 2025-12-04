(ns day4nt
  (:require [clojure.set :refer [union intersection difference]]))
(load-file "util.clj")


(def test-input
  "..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.")

(def real-input (util/read-input 4))

(defn parse-input
  [input]
  (let [grid (util/lines input)
        points (util/cartesian (range (count grid)) (range (count (grid 0))))]
    (->> points
         (filter #(= \@ (get-in grid %)))
         (set))))

(defn neighbours
  [point]
  (->> (util/cartesian (range -1 2) (range -1 2))
       (filter #(not= % [0 0]))
       (map (partial map + point))
       set))

(defn accessible?
  [points point]
  (< (count (intersection points (neighbours point))) 4))

(defn accessible-papers
  [points]
  (set (filter (partial accessible? points) points)))

(defn part-1
  [input]
  (->> input
       parse-input
       accessible-papers
       count))

(part-1 real-input)

(defn count-removable
  ([points] (count-removable points 0))
  ([points removed]
   (let [to-remove (accessible-papers points)]
     (if (empty? to-remove)
       removed
       (recur (difference points to-remove) (+ removed (count to-remove)))))))

(defn part-2
  [input]
  (->> input
       parse-input
       count-removable))

(time (part-2 real-input))
