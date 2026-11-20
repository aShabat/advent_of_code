┌─╴AOC
  Session ← ↘¯1⍜(&fo "../.session"|&rs∞)

  Get ← (
    ⊙Session
    $$ GET _ HTTP/1.0
    $$ Host: adventofcode.com
    $$ Connection: close
    $$ Cookie: session=_
    $$ 
    $$ 
    ⍜(&tlsc "adventofcode.com:443"|&rs∞˜⊸&w)
  )

  Post ← (
    ⊙(⊸⧻)
    ⊙Session
    $$ POST _ HTTP/1.1
    $$ Host: adventofcode.com
    $$ Connection: close
    $$ User-Agent: uiua
    $$ Content-Type: application/x-www-form-urlencoded
    $$ Cookie: session=_
    $$ Content-Length: _
    $$ 
    $$ _
    ⍜(&tlsc "adventofcode.com:443"|&rs∞˜⊸&w)
  )

  ResponseStatus ← ⋕ ⊙◌◌°$"_ _ _"
  ResponseBody   ← ◌°$"_\r\n\r\n_"
└─╴

SendAnswer ← (
  # day level answer
  AOC.ResponseStatus AOC.Post⊃($"/2026/day/_/answer"|$"level=_&answer=_"⋅∘)
)

LoadInput ← ⍜(&fo|&rs∞)$"static/_/input"

ParseRows! ← ⌅(≡(^0°□) ⊜□⊸(↥⊃(≠@\n|≠@\r))|/($"_\n_"∩°□)⍚°^0)
