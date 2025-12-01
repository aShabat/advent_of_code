(ns aoc-2025-1
  (:require [clojure.string :refer [split]]
            [clojure.math :refer [floor ceil]]))
(load-file "util.clj")

(def test-input "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82")

(def input (slurp "./static/1/input.txt"))

(defn parse-command [command]
  (let [matches (re-find #"(.)(\d+)" command)
        direction (get matches 1)
        length (Integer. (get matches 2))]
    {:dir direction :len length}))

(defn rotate [start {dir :dir len :len}]
  (mod ((case dir "R" + "L" -) start len) 100))

(defn reducer-1 [[start count] command]
  (let [end (rotate start command)]
    [end (if (= end 0) (inc count) count)]))

(defn solve-1 [input]
  (let [commands (map parse-command (split input #"\n"))
        start 50] (get  (reduce reducer-1 [start 0] commands) 1)))

(util/aoc-send-answer 2025 1 1 (solve-1 input))

(defn rotate-count-0 [start {dir :dir len :len}]
  (case dir
    "R" (let [pseudo-end (+ start len)
              count (inc (int (- (floor (/ pseudo-end 100)) (ceil (/ (inc start) 100)))))
              end (mod pseudo-end 100)] [end count])
    "L" (let [pseudo-end (- start len)
              count (inc (int (- (floor (/ (dec start) 100)) (ceil (/ pseudo-end 100)))))
              end (mod pseudo-end 100)] [end count])))

(defn reducer-2 [[start count] command]
  (let [[end add-count] (rotate-count-0 start command)] [end (+ count add-count)]))

(defn solve-2 [input]
  (let [commands (map parse-command (clojure.string/split input #"\n"))
        start 50
        end (reduce reducer-2 [start 0] commands)]
    (get end 1)))

(util/aoc-send-answer 2025 1 2 (solve-2 input))
