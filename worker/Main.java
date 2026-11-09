public class Main {
    public static void main(String[] args) throws InterruptedException {
        String name = System.getenv().getOrDefault("WORKER_NAME", "idp-test-app-multiple-worker");
        System.out.println(name + " started");
        long i = 0;
        while (true) {
            System.out.println(name + " heartbeat " + (i++));
            Thread.sleep(10000);
        }
    }
}
