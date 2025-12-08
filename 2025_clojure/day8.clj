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
       (keep-indexed (fn [index line]
                       {:coordinates (map parse-long (split line #",")),
                        :parent index,
                        :connections #{}}))
       vec))
(def test-state (parse-input test-input))

(defn dist [left right] (sqrt (reduce + (map (comp #(pow % 2) -) left right))))

(defn pairs [n] (mapcat (fn [k] (map #(list % k) (range k))) (range n)))

(defn light-connect
  [lights [left right]]
  (-> lights
      (update-in [left :connections] conj right)
      (update-in [right :connections] conj left)))

(defn lights-groups
  ([lights]
   (reduce (fn [[seen groups] index]
             (if-not (seen index)
               (let [[seen group] (lights-groups lights index seen)]
                 [seen (conj groups group)])
               [seen groups]))
           [#{} []]
           (range (count lights))))
  ([lights index seen]
   (if (seen index)
     [seen #{}]
     (let [seen (conj seen index)]
       (reduce (fn [[seen group] index']
                 (let [[seen group'] (lights-groups lights index' seen)]
                   [seen (into group group')]))
               [seen #{index}]
               (get-in lights [index :connections]))))))

(defn part-1
  [input connections]
  (let [lights (parse-input input)
        queue (->> (count lights)
                   pairs
                   (sort-by #(apply dist (map (comp :coordinates lights) %))))
        lights (reduce light-connect lights (take connections queue))]
    (->> (lights-groups lights)
         (second)
         (map count)
         sort
         reverse
         (take 3)
         (reduce *))))

(part-1 test-input 10)
(util/aoc-send-answer 8 1 (part-1 real-input 1000))
