'use client'

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { verifyEmail } from '@/utils/api';

export default function VerifyToken({ params }: { params: { token: string } }) {
  const [message, setMessage] = useState('Verifying your email...');
  const router = useRouter();

  useEffect(() => {
    const verify = async () => {
      try {
        await verifyEmail(params.token);
        setMessage('Email verified successfully. Redirecting to dashboard...');
        setTimeout(() => router.push('/dashboard'), 3000);
      } catch (error) {
        setMessage('Email verification failed. Please try again or contact support.');
      }
    };

    verify();
  }, [params.token, router]);

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">Email Verification</h2>
        <p className="mt-2 text-center text-sm text-gray-600">{message}</p>
      </div>
    </div>
  );
}

