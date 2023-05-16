 1. 특정 브랜치의 커밋 구조 확인하기

```
git log master
git log master --pretty=oneline
git log master --pretty=oneline --abbrev-commit
git log master --pretty=format:'%h %s' --graph
```

 2. 특정 브랜치 위치 확인하기

```
git rev-parse topic1
```

 3. 브랜치와 브랜치 사이에 있는 커밋들 확인하기

```
git log HEAD..origin/master
```

 4. 단순하게

```
git log --oneline --branches --graph
```