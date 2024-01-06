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
            position: 'top',
        }
    },
    interaction: {
        intersect: false,
        mode: 'index',
    },
    scales: {
        x: {
            border: {
                display: true
            },
            grid: {
                display: true,
                drawOnChartArea: true,
                drawTicks: true,
                color: 'rgba(93,93,93,0.2)',
            }
        },
        y: {
            grid: {
                display: true,
                drawOnChartArea: true,
                drawTicks: true,
                color: 'rgba(93,93,93,0.2)',
            }
        }
    }
}

const LineChart: React.FC<LineChartProps> = ({ chartData }) => {
    return <Line data={chartData} options={options} />;
};

export default LineChart;