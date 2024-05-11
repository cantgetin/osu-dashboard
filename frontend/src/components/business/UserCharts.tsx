import LineChart, {generateOptions} from "./LineChart.tsx";
import {convertDataToDayMonth} from "../../utils/utils.ts";
import {useEffect, useState} from "react";

interface LineChartProps {
    data: UserStatsDataset[];
    asSlideshow?: boolean;
    className?: string;
}

const charts: { property: keyof UserStatsDataset, name: string, color: string }[] = [
    {property: 'play_count', name: 'Plays', color: '#86EFAC'},
    {property: 'favourite_count', name: 'Favourites', color: '#FF5DBD'},
    {property: 'comments_count', name: 'Comments', color: '#f87171'}
];

const generateUserChartData = (
    userData: UserStatsDataset[],
    property: keyof UserStatsDataset,
    name: string,
    color: string
) => {
    const updatedData = userData.map((obj) => ({...obj, timestamp: convertDataToDayMonth(obj.timestamp)}));
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

const UserCharts = ({data, asSlideshow, className}: LineChartProps) => {
    const [currentIndex, setCurrentIndex] = useState(0);
    const intervalInSeconds = 3;

    const renderChart = (chartIndex: number) => {
        const chart = charts[chartIndex];
        return (<LineChart
            data={generateUserChartData(data, chart.property, chart.name, chart.color)}
            options={generateOptions(chart.name)}
        />);
    };

    useEffect(() => {
        const intervalId = setInterval(() => {
            setCurrentIndex((prevIndex) => (prevIndex + 1) % charts.length);
        }, intervalInSeconds * 1000);

        return () => clearInterval(intervalId);
    }, [charts.length, intervalInSeconds]);

    return (
        <>
            {data.length > 0 && (
                <div className={`flex bg-zinc-900 rounded-lg box-border ${className}`}>
                    {asSlideshow ? renderChart(currentIndex) : (
                        <div className="grid gap-4 grid-cols-1 w-full">
                            {charts.map((_, index) => (
                                <div key={index}>{renderChart(index)}</div>
                            ))}
                        </div>
                    )}
                </div>
            )}
        </>
    );
};

export default UserCharts;
