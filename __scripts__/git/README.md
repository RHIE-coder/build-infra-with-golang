# 인증서 저장하기

 - https://git-scm.com/book/ko/v2/Git-%EB%8F%84%EA%B5%AC-Credential-%EC%A0%80%EC%9E%A5%EC%86%8C

```
git config --global credential.helper cache # 15분까지 유지
git config --global credential.helper store # 디스크에 인증서 정보 저장
cat ~/.git-credentials                      # 인증서 확인하기
```

# Commit 에디터 변경하기

```
git config --global core.editor "vim"
```