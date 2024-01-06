import LineChart from "./LineChart.tsx";

interface LineChartProps {
    data: UserStatsDataset[];
}

const ChartsSummary = (props: LineChartProps) => {

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
        <div className="flex gap-3 bg-zinc-900 rounded-lg py-4 px-2 box-border w-full">
            <div className="w-1/2">
                <LineChart chartData={generateSingleChartData(props.data, 'play_count', '#86EFAC')}/>
            </div>
            {/*<div className="w-1/3">*/}
            {/*    <LineChart chartData={generateSingleChartData(props.data, 'map_count', '#64bbff')}/>*/}
            {/*</div>*/}
            <div className="w-1/2">
                <LineChart
                    chartData={generateSingleChartData(props.data, 'favourite_count', '#ff5dbd')}/>
            </div>
        </div>
    );
};

export default ChartsSummary;