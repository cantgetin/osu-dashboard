import React from 'react';
import { Line } from "react-chartjs-2";
import 'chart.js/auto';
import type {ChartOptions} from "chart.js";

interface LineChartProps {
    chartData: any;
}

const gridColor = 'rgba(115,115,115,0.2)';

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
                color: gridColor,
            }
        },
        y: {
            grid: {
                display: true,
                drawOnChartArea: true,
                drawTicks: true,
                color: gridColor,
            }
        }
    }
}

const LineChart: React.FC<LineChartProps> = ({ chartData }) => {
    return <Line data={chartData} options={options} />;
};

export default LineChart;