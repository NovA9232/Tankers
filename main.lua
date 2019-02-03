require "constants"

function love.load()
  SCREEN_W = 1000
  SCREEN_H = 800
  love.window.setMode(SCREEN_W, SCREEN_H)

  love.physics.setMeter(METER)
  world = love.physics.newWorld(0, 0, true)  -- true is for if bodies are allowed to sleep
  bodies = {players={}}
end

function love.draw()
  for _, j in pairs(bodies) do   -- Iterate each table in bodies.
    for x=1, #j do   -- For every body in that table, draw it.
      j[x]:draw()
    end
  end
end
