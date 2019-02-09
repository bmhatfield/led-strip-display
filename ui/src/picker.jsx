class Picker extends React.Component { // eslint-disable-line no-unused-vars

    interpolateColor(color1, color2, factor) {
        if (arguments.length < 3) {
            factor = 0.5;
        }

        let result = color1.slice();
        for (let i = 0; i < 3; i++) {
            result[i] = Math.round(result[i] + factor * (color2[i] - color1[i]));
        }
        return result;
    };

    interpolateColors(color1, color2, steps) {
        let stepFactor = 1 / (steps - 1);
        let interpolatedColorArray = [];

        color1 = color1.match(/\d+/g).map(Number);
        color2 = color2.match(/\d+/g).map(Number);

        for (let i = 0; i < steps; i++) {
            interpolatedColorArray.push(interpolateColor(color1, color2, stepFactor * i));
        }

        return interpolatedColorArray;
    }

    render() {
        // interpolateColors("rgb(94, 79, 162)", "rgb(247, 148, 89)", 5)

        return (
            <div className="picker" id="color-picker">
                <div className="color-group" id="b-w">
                    <p>Black and White</p>
                    <Color color="#000000" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FFFFFF" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#D68910" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#E2AC50" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FFEEBB" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="bold">
                    <p>Bold / Primary</p>
                    <Color color="#FF0000" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#00FF00" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#0000FF" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#6C3483" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FF8C00" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="pastel">
                    <p>Pastel</p>
                    <Color color="#83DDD6" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#8BEAAF" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#F7F396" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#F2C9C9" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#ACA7C4" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="winter">
                    <p>Winter</p>
                    <Color color="#FAE8DA" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#E6D2DE" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#C3B6E4" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#9B9EEB" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#5B7CEF" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="valentines">
                    <p>Valentines</p>
                    <Color color="#0D0045" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#61538A" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FFAACE" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FF6699" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#CC0033" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="st-patricks">
                    <p>St. Patricks</p>
                    <Color color="#93EB49" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FFFFFF" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FFC803" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FF8103" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#32593A" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="patriotic">
                    <p>Patriotic</p>
                    <Color color="#242038" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#0231EB" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#E8EDDF" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#A90818" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#540B0E" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="summer">
                    <p>Summer</p>
                    <Color color="#029DAF" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#E5D599" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FFC219" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#F07C19" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#E32551" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="halloween">
                    <p>Halloween</p>
                    <Color color="#FF9006" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FF6D10" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#2C0338" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#4D025C" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#800382" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="fall">
                    <p>Fall</p>
                    <Color color="#BF2A23" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#A6AD3C" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#F0CE4E" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#CF872E" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#8A211D" colorUpdate={this.props.broadcastHandler} />
                </div>

                <div className="color-group" id="christmas">
                    <p>Christmas</p>
                    <Color color="#E60F0D" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#FA860F" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#F8FF39" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#019406" colorUpdate={this.props.broadcastHandler} />
                    <Color color="#056099" colorUpdate={this.props.broadcastHandler} />
                </div>
            </div>
        );
    }
}
