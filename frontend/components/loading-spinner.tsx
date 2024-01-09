const LoadingSpinner = () => {
  return (
    <div className="fixed top-0 left-0 w-full h-full bg-[rgba(255,255,255,0.5)] flex justify-center items-center z-[999]">
      <div className="border-4 border-[#f3f3f3] border-t-primary rounded-full w-10 h-10 animate-spin"></div>
    </div>
  );
};

export default LoadingSpinner;
