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


(defn press
  [joltage switch]
  (reduce (fn [j n] (update j n dec)) (vec joltage) switch))

(def solve-2
  (memoize
   (fn [switches joltage]
     (let [initial-state [joltage switches 0 false]]
       (loop [[[joltage switches-left presses stopped?] & states']
              [initial-state]
              best-presses ##Inf]
         (cond (nil? joltage) best-presses
               (every? zero? joltage) (recur states' (min presses best-presses))
               (some neg? joltage) (recur states' best-presses)
               (and (every? even? joltage) (not stopped?))
               (let [presses' (+ presses
                                 (* 2
                                    (solve-2 switches (map #(/ % 2) joltage))))]
                 (recur (conj states' [joltage switches-left presses true])
                        (min presses' best-presses)))
               (seq switches-left)
               (let [[switch & switches-left'] switches-left
                     joltage' (press joltage switch)]
                 (recur (conj states'
                              [joltage switches-left' presses false]
                              [joltage' switches-left' (inc presses) false])
                        best-presses))
               :default (recur states' best-presses)))))))

(defn part-2
  [input]
  (let [data (parse-input input)]
    (reduce + (map (fn [d] (solve-2 (d :switches) (d :joltage))) data))))


(util/aoc-send-answer 10 2 (part-2 real-input))
