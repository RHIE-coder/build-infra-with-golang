// 구문 Syntax : 프로그래밍 문법
// 의미 semantics
// 아래 코드는 의미상 맞지 않는다.
const number = 'string';
const result = number * number;

// undefined: 개발자가 의도적으로 할당하기 위한 값이 아니라 엔진이 변수를 초기화할 때 사용함
// NaN: 은 자신과 일치하지 않는 유일한 값

function Falsy(v) {
    return !v;
}

function Truthy(v) {
    return !!v;
}
