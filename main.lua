require "constants"
require "tank"

function love.load()
  SCREEN_W = 1000
  SCREEN_H = 800
  love.window.setMode(SCREEN_W, SCREEN_H)

  bodies = {players={}}

  Tank(1, 200, 200)
end

function love.draw()
  for _, j in pairs(bodies) do   -- Iterate each table in bodies.
    for x=1, #j do   -- For every body in that table, draw it.
      j[x]:draw()
    end
  end
end
