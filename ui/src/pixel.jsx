class Pixel extends React.Component {
    render() {
        return (
            <div class="pixel" id={ `pixel-${this.props.side}-${this.props.id + 1}` }></div>
        )
    }
}