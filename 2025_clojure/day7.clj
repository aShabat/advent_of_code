(ns day7
  (:require [clojure.string :refer [trim]]
            [clojure.set :refer [difference intersection union]]))
(load-file "util.clj")

(def test-input
  ".......S.......
    ...............
    .......^.......
    ...............
    ......^.^......
    ...............
    .....^.^.^.....
    ...............
    ....^.^...^....
    ...............
    ...^.^...^.^...
    ...............
    ..^...^.....^..
    ...............
    .^.^.^.^.^...^.
    ...............")

(def real-input (util/read-input 7))

(defn parse-input
  [input]
  (let [lines (mapv trim (util/lines input))]
    [(first (keep-indexed (fn [i c] (when (= c \S) i)) (first lines)))
     (mapv #(set (keep-indexed (fn [i c] (when (= c \^) i)) %)) (rest lines))]))

(defn beam-descend-step
  [[split-count beam :as acc] splitters]
  (reduce (fn [[split-count beam] splitter]
            (if-not (contains? beam splitter)
              [split-count beam]
              [(inc split-count)
               (let [splitted (get beam splitter)]
                 (merge-with +
                             (dissoc beam splitter)
                             {(inc splitter) splitted,
                              (dec splitter) splitted}))]))
          acc
          splitters))

(defn beam-descend
  [beam splitters]
  (reduce beam-descend-step [0 {beam 1}] splitters))

(defn part-1
  [input]
  (let [[beam splitters] (parse-input input)]
    ((beam-descend beam splitters) 0)))

(defn part-2
  [input]
  (let [[beam splitters] (parse-input input)
        final-beam ((beam-descend beam splitters) 1)]
    (reduce + (vals final-beam))))
