class Pixel extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};
    }

    componentDidMount() {
        let options = {
            valueElement: null,
        };

        this.state.picker = new jscolor(`pixel-${this.props.side}-${this.props.id + 1}`, options);
    }

    render() {
        return (
            <div class={`pixel pixel-${this.props.side}`}
                 id={ `pixel-${this.props.side}-${this.props.id + 1}` }>
            </div>
        );
    }
}
