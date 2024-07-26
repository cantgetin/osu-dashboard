import {Doughnut} from "react-chartjs-2";
import {ChartOptions} from "chart.js";
import {useEffect, useState} from "react";

type UserDiagramsProps = {
    userId: number
    className: string
}

export const MakeChartHeightPlugin = (height: number) => {
    const ChartHeightPlugin = {
        id: 'customPlugin',
        beforeInit: function (chart: any) {
            // Get reference to the original fit function
            const originalFit = chart.legend.fit

            // Override the fit function
            chart.legend.fit = function fit() {
                // Bind scope in order to use `this` correctly inside it
                originalFit.bind(chart.legend)()
                this.height = height
            }
        }
    }

    return ChartHeightPlugin
}

const UserDiagrams = (props: UserDiagramsProps) => {
    const [userStats, setUserStats] = useState<UserStats | null>(null);

    useEffect(() => {
        fetch(`/api/user/statistic/${props.userId}`)
            .then(response => response.json())
            .then((data: UserStats) => setUserStats(data));
    }, [props.userId]);

    function makeOptions(key: string): ChartOptions<'doughnut'> {
        return {
            maintainAspectRatio: false,
            borderColor: 'black',
            cutout: '60%',
            plugins: {
                legend: {
                    display: true,
                    fullSize: false,
                    position: 'top',
                    labels: {
                        color: 'rgb(152,152,152)',
                        font: {
                            size: 12,
                            lineHeight: 10,
                            style: 'normal',
                            weight: 'normal',
                            family: 'Roboto',
                        },
                        boxWidth: 25,
                    }
                },
                title: {
                    display: true,
                    text: getNameFromKey(key),
                    position: 'top',
                    align: 'center',
                    font: {
                        size: 14,
                        style: 'normal',
                        weight: 'normal',
                        family: 'Roboto',
                    },
                    color: 'rgb(255,255,255)',
                },
            },
        };
    }

    interface NameMap {
        [key: string]: string;
    }

    function getNameFromKey(key: string): string {
        let nameMap: NameMap = {
            "most_popular_tags": "Tags",
            "most_popular_genres": "Genres",
            "most_popular_bpms": "BPM",
            "most_popular_starrates": "Starrate",
        };

        return nameMap[key];
    }

    interface Response {
        [key: string]: {
            [key: string]: number;
        };
    }

    const resp: Response = {
        "most_popular_tags": {
            "hip": 32,
            "hop": 32,
            "hyperpop": 21,
            "pop": 38,
            "rap": 58
        },
        "most_popular_genres": {
            "Electronic": 3,
            "Hip Hop": 10,
            "Rock": 3,
            "Unspecified": 45
        },
        "most_popular_bpms": {
            "140": 10,
            "150": 9,
            "160": 14,
            "170": 8,
            "200": 5
        },
        "most_popular_starrates": {
            "4": 15,
            "5": 24,
            "6": 21,
            "7": 4,
            "8": 3
        }
    }

    return (
        <div className={`p-4 bg-zinc-900 rounded-lg ${props.className}`}>
            <div className="grid grid-cols-2 gap-4">
                {userStats && Object.entries(resp).map(([key, value]) => (
                    <div key={key} className='h-80'>
                        <Doughnut
                            height="200px"
                            width="200px"
                            plugins={[MakeChartHeightPlugin(80)]}
                            data={{
                                labels: Object.keys(value),
                                datasets: [{
                                    data: Object.values(value),
                                    backgroundColor: ['#003f5c', '#58508d', '#bc5090', "#ff6361", "#ffa600"],
                                    borderRadius: 0,
                                    borderWidth: 0,
                                    label: getNameFromKey(key),
                                }],
                            }}
                            options={makeOptions(key)}
                            className="h-1/3"
                        />
                    </div>
                ))}
            </div>
        </div>
    );
};

export default UserDiagrams;