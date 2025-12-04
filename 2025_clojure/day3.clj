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
               bank)
       (dec n)))

(defn bank-joltage-stack
  [bank n]
  (loop [stack '()
         stack-length 0
         bank-tail bank
         bank-tail-length (count bank)]
    (cond (zero? bank-tail-length) (reduce #(+ (* 10 %1) %2) (reverse stack))
          (and (not (zero? stack-length))
               (> (first bank-tail) (first stack))
               (> (+ stack-length bank-tail-length) n))
          (recur (rest stack) (dec stack-length) bank-tail bank-tail-length)
          (< stack-length n) (recur (cons (first bank-tail) stack)
                                    (inc stack-length)
                                    (rest bank-tail)
                                    (dec bank-tail-length))
          :else
          (recur stack stack-length (rest bank-tail) (dec bank-tail-length)))))

(defn part-1
  [input]
  (->> input
       (parse-input)
       (map #(bank-joltage % 2))
       (reduce +)))

(defn part-2
  [input]
  (->> input
       (parse-input)
       (map #(bank-joltage % 12))
       (reduce +)))

(defn part-2-stack
  [input]
  (->> input
       (parse-input)
       (map #(bank-joltage-stack % 12))
       (reduce +)))
