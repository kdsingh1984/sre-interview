static int maxSubsetSum(int[] arr) {
        if(arr.length==1){
            return arr[0];
        }
        int[] dp = new int[arr.length];
        dp[0]=arr[0];
        dp[1]=Math.max(dp[0],arr[1]);
        for (int i = 2; i < arr.length; i++) {
            dp[i] = Math.max(Math.max(dp[i-1],arr[i]),dp[i-2]+arr[i]);
        }
        return dp[arr.length-1];
    }