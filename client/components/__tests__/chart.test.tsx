import React from 'react';
import { render, screen } from '@testing-library/react';
import { ChartTooltipContent } from '../ui/chart';

describe('ChartTooltipContent', () => {
  it('renders ChartTooltipContent component with no data', () => {
    const payload: any = [];

    render(<ChartTooltipContent payload={payload} />);

    expect(screen.queryByText('Item 2')).toBeFalsy()
    expect(screen.queryByText('200')).toBeFalsy()
  });
});