```go
now := time.Now().UnixNano()
rand.Seed(time.Now().UnixNano())
for i := 0; i < 20; i++ {
    // 랜덤하게 시간 부여(할당할 때마다 감소)
    for i := 0; i < len(dummySeed); i++ {
        if !filter(dummySeed[i]) {
            continue
        }
        randomNumber := int64(rand.Intn(999999999999999))
        descendingTime := now - randomNumber
        dummySeed[i].Datetime = time.Unix(0, descendingTime).Format(time.RFC3339)
        now = descendingTime
        dummyLog = append(dummyLog, dummySeed[i])
    }
}
```