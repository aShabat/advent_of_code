(ns day3)
(load-file "util.clj")

(def test-input
  "987654321111111
811111111111119
234234234234278
818181911112111")

(def real-input (util/read-input 3))

(defn parse-input
  [input]
  (->> input
       (util/lines)
       (map #(map (comp parse-long str first) (partition 1 %)))))

(defn bank-joltage
  [bank n]
  (nth (reduce (fn [acc digit]
               (map max acc (cons digit (map #(+ digit (* 10 %)) acc))))
          (repeat n 0)
          bank) n)

(bank-joltage [1 2 3] 2)
