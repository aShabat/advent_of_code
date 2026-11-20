Session ← ↘¯1⍜(&fo "../.session"|&rs∞)

AOCGet ← (
  $$ GET _ HTTP/1.1
  $$ Host: adventofcode.com
  $$ Connection: close
  $$ 
  $$ 
  ⍜(&tlsc "adventofcode.com:443"|&rs∞˜⊸&w)
)

AOCPost ← (
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
