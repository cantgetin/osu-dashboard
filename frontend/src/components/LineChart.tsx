import React from 'react';
import { Line } from "react-chartjs-2";
import 'chart.js/auto';
import type {ChartOptions} from "chart.js";

interface LineChartProps {
    chartData: any;
}

const options: ChartOptions<'line'> = {
    plugins:{
        legend: {
            display: false
        }
    },
    interaction: {
        intersect: false,
        mode: 'index',
    },
}

const LineChart: React.FC<LineChartProps> = ({ chartData }) => {
    return <Line data={chartData} options={options} />;
};

export default LineChart;