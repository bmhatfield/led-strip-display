class Pixel extends React.Component { // eslint-disable-line no-unused-vars
    constructor(props) {
        super(props);

        this.state = {
            selected: false,
            backgroundColor: props.bgcolor,
        };
    }

    updateSelection(event) {
        if ( event.buttons === 1 ) {
            this.setState({
                selected: !this.state.selected,
            });
        }
    }

    watchForESC(event) {
        // Clear all selections on Escape
        if (event.key === 'Escape') {
            this.setState({
                selected: false,
            });
        }

        // Select all on CMD+A
        if (event.key === 'a' && event.metaKey) {
            this.setState({
                selected: true,
            });
        }
    }

    style() {
        return {
            borderColor: this.state.selected ? '#35FCEE' : '#FFFFFF',
            backgroundColor: this.state.backgroundColor,
        };
    }

    componentDidUpdate(prevProps) {
        // Perform a number of checks to determine if we should actually set new state.

        // 1) Check if a new color came in via props.
        // This happens when a new strip is loaded into the Editor.
        let newBackgroundColor = prevProps.bgcolor !== this.props.bgcolor;

        // 2) Check if we've been sent a null selectionColor.
        // This happens after a color has been selected (onMouseUp on a Picker Color)
        let nullSelectionColor = this.props.selectionColor === null;

        // 3) Check if we've been sent a new selectionColor
        // This happens when a color is selected.
        let newSelectionColor = prevProps.selectionColor !== this.props.selectionColor;

        // 4) Check if we're selected and receiving a new color
        // This happens in the select -> pick case.
        let useSelectionColor = this.state.selected && newSelectionColor && !nullSelectionColor;

        // If we've got a new background color or a valid new color from the picker, update!
        if (newBackgroundColor || useSelectionColor) {
            this.setState({
                backgroundColor: newBackgroundColor ? this.props.bgcolor : this.props.selectionColor,
            });
        }
    }

    componentWillMount() {
        document.addEventListener('keydown', (event) => this.watchForESC(event));
    }

    render() {
        return (
            <label className={`pixel pixel-${this.props.side}`}
                 id={ `pixel-${this.props.side}-${this.props.id + 1}` }
                 style={ this.style() }
                 onMouseDown={ (event) => this.updateSelection(event) }
                 onMouseEnter={ (event) => this.updateSelection(event) }>
            </label>
        );
    }
}
