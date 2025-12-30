import {convertDataToDayMonth} from "../../../utils/time.ts";
import LineChart, {generateOptions} from "../stats/LineChart.tsx";


interface MapsetChartsProps {
    data: MapsetStatsDataset[]
    className?: string
}

const charts: { property: keyof MapsetStatsDataset, name: string, color: string }[] = [
    {property: 'play_count', name: 'Play count', color: '#86EFAC'},
    {property: 'favorite_count', name: 'Favourite count', color: '#FF5DBD'},
    {property: 'comments_count', name: 'Comment count', color: '#f87171'}
];

const generateMapChartData = (
    mapsetData: MapsetStatsDataset[],
    property: keyof MapsetStatsDataset,
    name: string,
    color: string
) => {
    const updatedData =
        mapsetData.map((obj) => ({...obj, timestamp: convertDataToDayMonth(obj.timestamp)}));
    const slicedData = updatedData.length >= 7 ? updatedData.slice(-7) : updatedData;
    return {
        labels: slicedData.map((data) => data.timestamp),
        datasets: [{
            data: slicedData.map((data) => data[property]),
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

const MapsetCharts = (props: MapsetChartsProps) => {

    const renderChart = (chartIndex: number) => {
        const chart = charts[chartIndex];
        return (<LineChart
            data={generateMapChartData(props.data, chart.property, chart.name, chart.color)}
            options={generateOptions(chart.name)}
        />);
    };

    return (
        <>
            {
                props.data != null ?
                    <div className={`flex bg-zinc-900 box-border w-full ${props.className}`}>
                        <div className="grid 2xl:grid-cols-2 lg:grid-cols-2 w-full gap-2">
                            {charts.map((_, index) => (
                                <div key={index}>{renderChart(index)}</div>
                            ))}
                        </div>
                    </div>
                    : null
            }
        </>
    );
};

export default MapsetCharts;