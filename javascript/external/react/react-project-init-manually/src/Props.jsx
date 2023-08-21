const PropsComponent = props => {
    return (
        <div>제이름은 {props.name}입니다.</div>
    )
}

PropsComponent.defaultProps = {
    name: 'rhie-coder',
}

export default PropsComponent;