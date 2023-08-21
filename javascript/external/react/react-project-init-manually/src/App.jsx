import PropTypes from 'prop-types'

const App = props => {
    return (
        <div>제이름은 {props.name}입니다.</div>
    )
}

App.defaultProps = {
    name: 'rhie-coder',
}

App.propTypes = {
    name: PropTypes.string
}

export default App;