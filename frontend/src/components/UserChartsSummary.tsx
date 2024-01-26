import LineChart from "./LineChart.tsx";
import {convertDataToDayMonth} from "../utils/utils.ts";

interface LineChartProps {
    data: UserStatsDataset[]
    onlyPlaycount?: boolean
}

const UserChartsSummary = (props: LineChartProps) => {

    function mapToChartData(data: UserStatsDataset[]): UserStatsDataset[] {
        const updatedArray: UserStatsDataset[] = [];
        data.forEach((obj) => updatedArray.push({...obj, timestamp: convertDataToDayMonth(obj.timestamp)}));
        return updatedArray
    }

    type ChartDataProperty = keyof UserStatsDataset;

    const generateSingleChartData = (
        userData: UserStatsDataset[],
        property: ChartDataProperty,
        name: string,
        color: string
    ) => {
        if (userData.length >= 7) {
            userData = userData.slice(-7)
        }

        return {
            labels: userData.map((data) => data.timestamp),
            datasets: [{
                data: userData.map((data) => data[property]),
                backgroundColor: [color],
                label: name,
                borderColor: [color],
                borderWidth: 2,
                pointStyle: 'circle',
                pointRadius: 3,
                pointHoverRadius: 6,
            }],
        };
    };

    const playCountChart = () => {
        return (
            <LineChart
                chartData={
                    generateSingleChartData(
                        mapToChartData(props.data),
                        'play_count',
                        'Play count',
                        '#86EFAC'
                    )
                }
            />
        )
    }

    const favouritesChart = () => {
        return (
            <LineChart
                chartData={
                    generateSingleChartData(
                        mapToChartData(props.data),
                        'favourite_count',
                        'Favourite count',
                        '#FF5DBD'
                    )
                }
            />
        )
    }

    return (
        <>
            {
                props.data.length > 0 ?
                    <div className="flex gap-3 bg-zinc-900 rounded-lg p-2 box-border w-full">
                        {props.onlyPlaycount ?
                            <div className="w-full">
                                {playCountChart()}
                            </div> :
                            <>
                                <div className="w-1/2">
                                    {playCountChart()}
                                </div>
                                <div className="w-1/2">
                                    {favouritesChart()}
                                </div>
                            </>
                        }
                    </div>
                    : null}
        </>
    );
};

export default UserChartsSummary;