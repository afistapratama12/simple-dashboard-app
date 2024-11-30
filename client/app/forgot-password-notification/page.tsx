export default function ForgotPasswordNotification() {
    return (
      <div className="min-h-screen bg-gray-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-md">
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">Check Your Email</h2>
          <p className="mt-2 text-center text-sm text-gray-600">
            We&apos;ve sent a password reset link to your email address. Please check your inbox and click the reset button to set a new password.
          </p>
        </div>
      </div>
    );
  }
  
  