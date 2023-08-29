import {useLocation} from 'react-router-dom';

const About = () => {
    const location = useLocation();

    return (
        <div>
            <h1>introduce</h1>
            <p>what is the react? good good</p>
            <p>querystring(pathname): {location.pathname}</p>
            <p>querystring(seacrh): {location.search}</p> {/* 맨 앞의 ? 문자를 포함한 쿼리스트링 값 */}
            <p>querystring(hash): {location.hash}</p> {/* 주소의 # 문자열 뒤의 값 */}
            <p>querystring(state): {location.state}</p> {/* 페이지로 이동할 때 임의로 넣을 수 있는 상태값 */}
            <p>querystring(key): {location.key}</p> {/* location 객체의 고유값 */}
        </div>
    )
}

export default About;