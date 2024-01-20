import LineChart from "./LineChart.tsx";
import {convertDataToDayMonth} from "../utils/utils.ts";

interface LineChartProps {
    data: UserStatsDataset[];
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
        color: string
    ) => {
        return {
            labels: userData.map((data) => data.timestamp),
            datasets: [{
                data: userData.map((data) => data[property]),
                backgroundColor: [color],
                label: property, // Add the label for the legend
                borderColor: [color],
                borderWidth: 2,
                pointStyle: 'circle',
                pointRadius: 3,
                pointHoverRadius: 6,
            }],
        };
    };

    return (
        <div className="flex gap-3 bg-zinc-900 rounded-lg p-2 box-border w-full">
            <div className="w-1/2">
                <LineChart chartData={generateSingleChartData(mapToChartData(props.data), 'play_count', '#86EFAC')}/>
            </div>
            <div className="w-1/2">
                <LineChart chartData={generateSingleChartData(mapToChartData(props.data), 'favourite_count', '#ff5dbd')}/>
            </div>
        </div>
    );
};

export default UserChartsSummary;