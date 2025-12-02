#!/usr/bin/clj -M
(load-file "util.clj")

(let [command (first *command-line-args*)]
  (case command
    "get" (do (apply util/aoc-get-exercise (rest *command-line-args*))
              (apply util/aoc-get-input (rest *command-line-args*))
              (prn "success!"))
    "answer" (print (apply util/aoc-send-answer (rest *command-line-args*)))
    (prn "wrong command")))
(shutdown-agents)
