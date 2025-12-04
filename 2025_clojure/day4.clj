(ns day4)
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
  (->> input
       (util/lines)
       (mapv (fn [line] (mapv #(if (= \@ %) 1 0) line)))))

(defn grid-data [grid] {:height (count grid), :width (count (get grid 1))})

(defn neighbours
  [position]
  (->> (util/cartesian (range -1 2) (range -1 2))
       (filter #(not= '(0 0) %))
       (map #(map + position %))))

(defn access?
  [position grid]
  (and (= 1 (get-in grid position))
       (->> (neighbours position)
            (map #(get-in grid % 0))
            (reduce +)
            (> 4))))

(access? '(0 2) (parse-input test-input))

(defn part-1
  [input]
  (let [grid (parse-input input)
        positions (util/cartesian (range (count grid))
                                  (range (count (grid 0))))]
    (->> (filter #(access? % grid) positions)
         (count))))

(part-1 test-input) ;should be equal 13
(util/aoc-send-answer 4 1 (part-1 real-input))

(defn remove-all-papers
  ([grid positions-to-check skipped-positions]
   (if (empty? positions-to-check)
     grid
     (let [position (first positions-to-check)]
       (cond (not (get-in grid position))
             (recur grid (rest positions-to-check) skipped-positions)
             (access? position grid) (recur (assoc-in grid position false)
                                            (concat (rest positions-to-check)
                                                    skipped-positions)
                                            '())
             :else (recur grid
                          (rest positions-to-check)
                          (cons position skipped-positions))))))
  ([grid]
   (remove-all-papers grid
                      (util/cartesian (range (count grid))
                                      (range (count (grid 0))))
                      '())))


