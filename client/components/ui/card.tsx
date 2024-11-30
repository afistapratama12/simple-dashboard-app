import React from 'react';

const Card = ({ children }: any) => <div className="card">{children}</div>;

const CardHeader = ({ children }: any) => <header className="card-header">{children}</header>;

const CardFooter = ({ children }: any) => <footer className="card-footer">{children}</footer>;

const CardTitle = ({ children }: any) => <h2 className="card-title text-black">{children}</h2>;

const CardDescription = ({ children }: any) => <p className="card-description text-black">{children}</p>;

const CardContent = ({ children }: any) => <div className="card-content">{children}</div>;

export { Card, CardHeader, CardFooter, CardTitle, CardDescription, CardContent }

