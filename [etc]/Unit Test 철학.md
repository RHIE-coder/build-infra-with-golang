# The Four Levels of Software Testing

 https://www.seguetech.com/the-four-levels-of-software-testing/

 - unit test
1. 함수 하나에서 모두 끝나는 것. input - output이 확실한 것.
2. 자체 오브젝트(class)에 한정된 테스트

 - integration test(실제 real data, not mocking data) -- blackbox, smoking
3. 모듈 간의 dependencies가 존재할 때
4. 다른 모듈의 상태를 의존함

<hr>

`3.`에 Mocking하려면 어떻게면

함수 대신에 모듈을 대신 받는게 아니라, 모듈에서 나온 데이터를 받게 하면 된다.

이 경우 가짜 데이터만 주면 됨.

이렇게 디자인하면 클래스가 간단해진다.(모두 적용은 안될 듯)

`4.` 다른 함수 및 모듈의 결과에서 나온 것을 받아서 사용하는 함수에 대한 테스트이면?

Receiver 역할을 하는 함수에 대해서만 mocking 데이터를 제공해서 테스트하면 된다.

만일 다 안되면?

함수 모두를 Mocking 하느냐? No 그냥 하지마라.




