import {Route, Routes} from 'react-router-dom';
import About from './About';
import Home from './Home';
import Profile from './Profile';

const App = () => {
    return (
        <Routes>
            <Route path="/" element={<Home></Home>}></Route>
            <Route path="/about" element={<About></About>}></Route>
            <Route path="/profiles/:username" element={<Profile></Profile>}></Route>
        </Routes>
    )
}

export default App;