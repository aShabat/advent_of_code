(ns day11
  (:require [clojure.string :refer [trim split]]))
(load-file "util.clj")

(def test-input
  "aaa: you hhh
    you: bbb ccc
    bbb: ddd eee
    ccc: ddd eee fff
    ddd: ggg
    eee: out
    fff: out
    ggg: out
    hhh: ccc fff iii
    iii: out")

(def real-input (util/read-input 11))

(defn parse-input
  [input]
  (->> (util/lines input)
       (map trim)
       (map (fn [line]
              (let [words (split line #" ")
                    in (first words)
                    in (subs in 0 (dec (count in)))
                    out (rest words)]
                {in out})))
       (apply conj)))

(parse-input test-input)

(defn generate-memo-dfs
  [graph]
  (let [dfs
        (fn [recur-dfs start end]
          (cond (= start end) 1
                (not (contains? graph start)) 0
                :default
                (reduce + 0 (map #(recur-dfs recur-dfs % end) (graph start)))))
        memo-dfs (memoize dfs)]
    (partial memo-dfs memo-dfs)))

(defn part-1
  [input]
  (let [graph (parse-input input)
        dfs (generate-memo-dfs graph)]
    (dfs "you" "out")))

(util/aoc-send-answer 11 1 (part-1 real-input))

(defn part-2
  [input]
  (let [graph (parse-input input)
        dfs (generate-memo-dfs graph)]
    (+ (* (dfs "svr" "fft") (dfs "fft" "dac") (dfs "dac" "out"))
       (* (dfs "svr" "dac") (dfs "dac" "fft") (dfs "fft" "out")))))

; (util/aoc-send-answer 11 2 (part-2 real-input))
