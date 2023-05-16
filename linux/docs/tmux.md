# # --- Session --- #

## [생성]

```sh
$ tmux
$ tmux new -s <session_name>
$ tmux new-session -s <session_name>
```

## [조회]

```sh
$ tmux ls
```

## [연결]

### - detach

```
[CTRL] + b, d
```

### - attach

```sh
$ tmux attach -t <session_number || session_name>
```

## [수정]

```
[CTRL] + b, $
```

## [삭제]

```sh
tmux kill-session -t <session_name>
```

<hr><br><br><br>

# # --- Window --- #

<hr><br><br><br>

# # --- Pane --- #
