import {Doughnut} from "react-chartjs-2";
import {ChartOptions} from "chart.js";
import {useAppSelector} from "../../../store/hooks.ts";
import {LoadingState} from "../../../interfaces/LoadingState.ts";
import {selectUserStats, selectUserStatsLoading,} from "../../../store/userStatsSlice.ts";
import LoadingSpinner from "../../ui/LoadingSpinner.tsx";
import {useFetchUserStatsOnce} from "../../../hooks/useFetchUserStats.ts";

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
    const userStats = useAppSelector<UserStatistics | null>(selectUserStats)
    const userStatsLoaded = useAppSelector<LoadingState>(selectUserStatsLoading)

    useFetchUserStatsOnce(props.userId.toString())

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
        const nameMap: NameMap = {
            "most_popular_tags": "Tags",
            "most_popular_genres": "Genres",
            "most_popular_bpms": "BPM",
            "most_popular_starrates": "Starrate",
        };

        return nameMap[key];
    }

    function ensureValueHasData(value: UserStatisticUnit): UserStatisticUnit {
        if (value && typeof value === 'object' && !Array.isArray(value)) {
            if (Object.keys(value).length === 0) {
                return {Unspecified: 100};
            }
        }
        return value;
    }

    return (
        <>
            {
                userStatsLoaded == LoadingState.Succeeded ?
                    <div className={`p-2 md:p-4 bg-zinc-900 rounded-lg ${props.className}`}>
                        <div className="grid grid-cols-1 sm:grid-cols-2 gap-2 md:gap-4">
                            {userStats && Object.entries(userStats).map(([key, value]) => {
                                const modifiedValue = ensureValueHasData(value);
                                return getNameFromKey(key) != undefined ? (
                                    <div key={key} className='h-48 md:h-80'>
                                        <Doughnut
                                            height="100%"
                                            width="100%"
                                            plugins={[MakeChartHeightPlugin(80)]}
                                            data={{
                                                labels: Object.keys(modifiedValue).map(item => item === "" ? "Unspecified" : item),
                                                datasets: [{
                                                    data: Object.values(modifiedValue),
                                                    backgroundColor: ['#003f5c', '#58508d', '#34ab9c', "#ff6361", '#983a73'],
                                                    borderRadius: 0,
                                                    borderWidth: 0,
                                                    label: getNameFromKey(key),
                                                }],
                                            }}
                                            options={makeOptions(key)}
                                        />
                                    </div>
                                ) : null;
                            })}
                        </div>
                    </div>
                    : <LoadingSpinner/>
            }
        </>
    );
};

export default UserDiagrams;