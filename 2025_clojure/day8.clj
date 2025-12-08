(ns day8
  (:require [clojure.string :refer [split]]
            [clojure.math :refer [sqrt pow]]))
(load-file "util.clj")

(def test-input
  "162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689")

(def real-input (util/read-input 8))

(defn parse-input
  [input]
  (->> input
       util/lines
       (keep-indexed (fn [index line] (map parse-long (split line #","))))
       vec))

(defn dist [left right] (sqrt (reduce + (map (comp #(pow % 2) -) left right))))

(defn pairs [n] (mapcat (fn [k] (map #(list % k) (range k))) (range n)))

(defn find-cirquit [cirquits index] (first (filter #(% index) cirquits)))

(defn connect
  [cirquits [left right]]
  (let [left-cirquit (find-cirquit cirquits left)
        right-cirquit (find-cirquit cirquits right)]
    (if (= left-cirquit right-cirquit)
      cirquits
      (conj (disj cirquits left-cirquit right-cirquit)
            (into left-cirquit right-cirquit)))))

(defn part-1
  [input connections]
  (let [lights (parse-input input)
        queue (->> (count lights)
                   pairs
                   (sort-by #(apply dist (map lights %)))
                   (take connections))
        cirquits (set (map #(set (list %)) (range (count lights))))]
    (prn queue)
    (->> (reduce connect cirquits queue)
         (map count)
         sort
         reverse
         (take 3)
         (reduce *))))

; (part-1 test-input 10)
; (util/aoc-send-answer 8 1 (part-1 real-input 1000))

(defn part-2
  [input]
  (let [lights (parse-input input)
        queue (->> (count lights)
                   pairs
                   (sort-by #(apply dist (map lights %))))
        cirquits (set (map #(set (list %)) (range (count lights))))]
    (reduce (fn [cirquits [left right :as pair]]
              (let [cirquits (connect cirquits pair)]
                (if (= 1 (count cirquits))
                  (reduced (* (first left) (first right)))
                  cirquits)))
            cirquits
            queue)))
