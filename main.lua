require "constants"
require "tank"
require "vec2"

lg = love.graphics

function love.load()
  SCREEN_W = 1000
  SCREEN_H = 800
  love.window.setMode(SCREEN_W, SCREEN_H)

  entities = {players={}}

  Tank(1, 500, 400)
end

function love.draw()
  for _, j in pairs(entities) do   -- Iterate each table in bodies.
    for x=1, #j do   -- For every body in that table, draw it.
      j[x]:draw()
    end
  end

  lg.print(love.timer.getFPS(), 10, 10)
end

function love.update(dt)
  for _, j in pairs(entities) do
    for x=1, #j do
      j[x]:update(dt)
    end
  end
end
