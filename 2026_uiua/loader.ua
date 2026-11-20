~ "util.ua" ~ AOC

≡(&fmd$"static/_")+1⇡12

+1⇡12
≡(
  AOC.Get⊸$"/2026/day/_"
  ⊃(=200AOC.ResponseStatus|AOC.ResponseBody)
  ⍥(
    ⊸∘⍜(⊙&fc|&w) ⊙⊸$"static/_/problem.html"
    ◌&runi {"pandoc" ⊙"-o" $"static/_/problem.html" ⟜$"static/_/problem.md"}
    AOC.ResponseBody AOC.Get⊸$"/2026/day/_/input"
    ⍜(⊙&fc|&w)⊙⊸$"static/_/input"
    &p
  )
)
