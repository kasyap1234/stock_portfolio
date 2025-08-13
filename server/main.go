package main 


func main(){
	   cfg := config.GetConfig()
    log.Printf("Server starting on port %d", cfg.Server.Port)
    
    e := echo.New()
    
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        if err := e.Start(":" + config.GetString("server.port")); err != nil {
            log.Printf("shutting down the server: %v", err)
        }
    }()
    
    <-quit
    
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    log.Println("shutting down")
    if err := e.Shutdown(shutdownCtx); err != nil {
        log.Fatal("Server failed to shutdown gracefully")
    }
    
    log.Print("server stopped")
}
}