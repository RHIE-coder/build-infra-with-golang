import {Link} from 'react-router-dom';

const Home = () => {
    return (
        <div>
            <h1>Home</h1>
            <p>the main page</p>
            <ul>
                <li>
                    <Link to="/about">ABOUT</Link> 
                </li>
                <li>
                    <Link to="/profiles/memberA">memberA</Link> 
                </li>
                <li>
                    <Link to="/profiles/memberB">memberB</Link> 
                </li>
                <li>
                    <Link to="/profiles/membera">membera</Link> 
                </li>
            </ul>
        </div>
    )
}

export default Home;