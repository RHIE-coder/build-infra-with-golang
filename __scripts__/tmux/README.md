---
Sessions:
  Window 0:
    - Pane
    - Pane
    - Pane
  Window 1:
    - Pane
    - Pane
    - Pane
---

# tmux

 - 리눅스 터미널 분할 툴

## [ Install ]

```sh
sudo apt install tmux
```
## [ Structure ]

### - TMUX Server

 - Kill

```sh
tmux kill-server
tmux kill-session -t <session_name>
```

### - Session

```sh
tmux                            # 세션 생성 및 진입
tmux new -s <session_name>      # 세션 생성 및 진입(이름 설정)
[PREFIX] + d                    # 세션 나가기(detach)
CTRL + d                        # 세션 종료하고 나가기
tmux ls                         # 세션 리스트
tmux a(ttach) -t <session_name> # 특정 세션으로 들어가기
tmux a #                        # 직전 detach된 세션으로 들어가기
[PREFIX] + $                    # 현재 세션 이름 바꾸기
```

### - Window / Pane

```
[PREFIX] + [                // Window 스크롤(Page Up / Page Down)
[PREFIX] + %                // 수직 분할
[PREFIX] + "                // 수평 분할
[PREFIX] + 방향키           // Pane간의 이동
[PREFIX] + x                // Pane kill
[PREFIX] + &                // window kill
[PREFIX] + CTRL + 방향키    // Pane 크기 변경
```


## [ MORE ]

```sh
C-z         Suspend the tmux client.
!           Break the current pane out of the window.
(           Switch the attached client to the previous session.
)           Switch the attached client to the next session.
,           Rename the current window.
0 to 9      Select windows 0 to 9.
:           Enter the tmux command prompt.
;           Move to the previously active pane.
D           Choose a client to detach.
L           Switch the attached client back to the last session.
[           Enter copy mode to copy text or view the history.
]           Paste the most recently copied buffer of text.
c           Create a new window.
i           Display some information about the current window.
l           Move to the previously selected window.
n           Change to the next window.
o           Select the next pane in the current window.
p           Change to the previous window.
q           Briefly display pane indexes.
m           Mark the current pane (see select-pane -m).
M           Clear the marked pane.
s           Select a new session for the attached client interactively.
t           Show the time.
w           Choose the current window interactively.
z           Toggle zoom state of the current pane.
{           Swap the current pane with the previous pane.
}           Swap the current pane with the next pane.

C-Up, C-Down
C-Left, C-Right
            Resize the current pane in steps of one cell.
```