class Picker extends React.Component { // eslint-disable-line no-unused-vars
    render() {
        return (
            <div className="picker" id="color-picker">
                <Color color="#000000" colorUpdate={this.props.broadcastHandler} />
                <Color color="#FF0000" colorUpdate={this.props.broadcastHandler} />
                <Color color="#00FF00" colorUpdate={this.props.broadcastHandler} />
                <Color color="#0000FF" colorUpdate={this.props.broadcastHandler} />
                <Color color="#6C3483" colorUpdate={this.props.broadcastHandler} />
                <Color color="#D68910" colorUpdate={this.props.broadcastHandler} />
                <Color color="#E2AC50" colorUpdate={this.props.broadcastHandler} />
                <Color color="#FFEEBB" colorUpdate={this.props.broadcastHandler} />
                <Color color="#FFFFFF" colorUpdate={this.props.broadcastHandler} />
            </div>
        );
    }
}
