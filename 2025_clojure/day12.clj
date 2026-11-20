(ns day12
  (:require [clojure.string :refer [trim split]]))
(load-file "util.clj")

(def test-input
  "0:
    ###
    ##.
    ##.

    1:
    ###
    ##.
    .##

    2:
    .##
    ###
    ##.

    3:
    ##.
    ###
    ##.

    4:
    ###
    #..
    ###

    5:
    ###
    .#.
    ###

    4x4: 0 0 0 0 2 0
    12x5: 1 0 1 0 2 2
    12x5: 1 0 1 0 3 2")

(def real-input (util/read-input 12))

(defn parse-input
  [input]
  (let [ps (split input #"\n\n")
        presents (butlast ps)
        trees (last ps)]
    [(map (fn [present] (count (re-seq #"#" present))) presents)
     (map (fn [line]
            (let [words (split (trim line) #" ")
                  size (first words)
                  counts (rest words)]
              [(map parse-long (split size #"x|:")) (map parse-long counts)]))
          (util/lines trees))]))

(parse-input real-input)

(defn part-1-very-stupid
  [input]
  (let [[presents trees] (parse-input input)]
    (count (filter (fn [[[width height] counts]]
                     (>= (* width height) (reduce + (map * presents counts))))
                   trees))))

(util/aoc-send-answer 12 1 (part-1-very-stupid real-input))
