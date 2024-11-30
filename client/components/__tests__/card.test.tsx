

import React from 'react';
import { render } from '@testing-library/react';
import { Card, CardHeader, CardFooter, CardTitle, CardDescription, CardContent } from '../ui/card';

describe('Card components', () => {
    it('renders Card component with children', () => {
        const { getByText } = render(<Card>Card Content</Card>);
        expect(getByText('Card Content')).toBeTruthy();
    })

    it('renders CardHeader component with children', () => {
        const { getByText } = render(<CardHeader>Header Content</CardHeader>);
        expect(getByText('Header Content')).toBeTruthy();
    });

    it('renders CardFooter component with children', () => {
        const { getByText } = render(<CardFooter>Footer Content</CardFooter>);
        expect(getByText('Footer Content')).toBeTruthy();
    });

    it('renders CardTitle component with children', () => {
        const { getByText } = render(<CardTitle>Title Content</CardTitle>);
        expect(getByText('Title Content')).toBeTruthy();
    });

    it('renders CardDescription component with children', () => {
        const { getByText } = render(<CardDescription>Description Content</CardDescription>);
        expect(getByText('Description Content')).toBeTruthy();
    });

    it('renders CardContent component with children', () => {
        const { getByText } = render(<CardContent>Content</CardContent>);
        expect(getByText('Content')).toBeTruthy();
    });
});